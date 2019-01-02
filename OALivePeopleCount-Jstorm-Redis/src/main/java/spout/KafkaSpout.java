package spout;

/**
 * Created by hzcortex on 16-11-15.
 */
import java.util.Collection;
import java.util.Map;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import spout.KafkaConsumerUtil.KafkaMessageID;
import spout.KafkaConsumerUtil.PartitionConsumer;
import spout.KafkaConsumerUtil.PartitionConsumer.EmitState;
import spout.KafkaConsumerUtil.PartitionCoordinator;

import backtype.storm.spout.SpoutOutputCollector;
import backtype.storm.task.TopologyContext;
import backtype.storm.topology.IRichSpout;
import backtype.storm.topology.OutputFieldsDeclarer;
import backtype.storm.tuple.Fields;

import config.KafkaSpoutConfig;
/*import com.xueersi.jstorm.config.KafkaSpoutConfig;*/
import spout.KafkaConsumerUtil.ZKState;

public class KafkaSpout implements IRichSpout {
    private static final long serialVersionUID = 1L;
    /*private static Logger LOG = LoggerFactory.getLogger(KafkaSpout.class);*/

    protected SpoutOutputCollector collector;

    private long lastUpdateMs;
    private PartitionCoordinator coordinator;


    private KafkaSpoutConfig config;

    private ZKState zkState;


    public KafkaSpout(KafkaSpoutConfig config) {
        this.config = config;
    }


    public void open(Map conf, TopologyContext context, SpoutOutputCollector collector) {
        try {
            this.collector = collector;
            if (this.config == null) {
                config = new KafkaSpoutConfig();
                config.configure(conf);
            }
            zkState = new ZKState(conf, config);
            coordinator = new PartitionCoordinator(conf, config, context, zkState);
            lastUpdateMs = System.currentTimeMillis();
        } catch (Exception ex) {
            ex.printStackTrace();
        }
    }

    @Override
    public void close() {
        // TODO Auto-generated method stub
        zkState.close();
    }

    @Override
    public void activate() {
        // TODO Auto-generated method stub

    }

    @Override
    public void deactivate() {
        // TODO Auto-generated method stub

    }

    @Override
    public void nextTuple() {
        try {
            Collection<PartitionConsumer> partitionConsumers = coordinator.getPartitionConsumers();
            boolean isAllSleeping = true;
            for (PartitionConsumer consumer : partitionConsumers) {
                if (!consumer.isSleepingConsumer()) {
                    isAllSleeping = false;
                    EmitState state = consumer.emit(collector);
                    /*LOG.debug("====== partition " + consumer.getPartition() + " emit message state is " + state);*/
                   /* String timeStamp = new SimpleDateFormat("yyyyMMdd_HHmmss").format(Calendar.getInstance().getTime());
                    LOG.info("read data from kafka in" + timeStamp);*/
                }

            }
            if (isAllSleeping) {
                try {
                    Thread.sleep(100);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            }
            long now = System.currentTimeMillis();
            if ((now - lastUpdateMs) > config.offsetUpdateIntervalMs) {
                commitState();
            }
        } catch (Exception ex) {
            ex.printStackTrace();
        }

    }

    public void commitState() {
        lastUpdateMs = System.currentTimeMillis();
        for (PartitionConsumer consumer : coordinator.getPartitionConsumers()) {
            consumer.commitState();
        }

    }

    @Override
    public void ack(Object msgId) {
        KafkaMessageID messageId = (KafkaMessageID) msgId;
        PartitionConsumer consumer = coordinator.getConsumer(messageId.getPartition());
        consumer.ack(messageId.getOffset());
    }

    @Override
    public void fail(Object msgId) {
        KafkaMessageID messageId = (KafkaMessageID) msgId;
        PartitionConsumer consumer = coordinator.getConsumer(messageId.getPartition());
        consumer.fail(messageId.getOffset());
    }

    @Override
    public void declareOutputFields(OutputFieldsDeclarer declarer) {
        declarer.declare(new Fields("logfile"));

    }

    @Override
    public Map<String, Object> getComponentConfiguration() {
        return null;
    }


}

