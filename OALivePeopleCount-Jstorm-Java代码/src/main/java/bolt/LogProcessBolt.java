package bolt;

/**
 * Created by hzcortex on 16-11-15.
 * deal the logic with Jedis
 */

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.Executor;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import backtype.storm.topology.BasicOutputCollector;
import backtype.storm.task.TopologyContext;
import backtype.storm.topology.OutputFieldsDeclarer;
import backtype.storm.topology.base.BaseBasicBolt;
import backtype.storm.tuple.Tuple;
import backtype.storm.tuple.Fields;

import config.MainConfig;
import org.apache.commons.pool2.impl.GenericObjectPoolConfig;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.exceptions.JedisConnectionException;
import redis.clients.jedis.exceptions.JedisDataException;
import util.JedisAPI;

class InsertRedisThread implements Runnable {
    private Jedis jedis;
    private String key_;
    private String stuId;

    public InsertRedisThread(Jedis jedis, String key_, String stuId) {
        this.jedis = jedis;
        this.key_ = key_;
        this.stuId = stuId;
    }

    @Override
    public void run() {
        try {
            System.out.println("增加使用的线程为： " + Thread.currentThread().getName());
            jedis.select(MainConfig.redisDatabaseIndex);
            jedis.sadd(key_, stuId);
            jedis.close();
        } catch (JedisDataException ex) {// multiple catch is not supported before java7
            ex.printStackTrace();
        } catch (NullPointerException ex) {
            ex.printStackTrace();
        }

    }
}

class RemoveRedisThread implements Runnable {
    private Jedis jedis;
    private String key_;
    private String stuId;

    public RemoveRedisThread(Jedis jedis, String key_, String stuId) {
        this.jedis = jedis;
        this.key_ = key_;
        this.stuId = stuId;
    }

    @Override
    public void run() {
        try {
            System.out.println("移除使用的线程为： " + Thread.currentThread().getName());
            jedis.select(MainConfig.redisDatabaseIndex);
            jedis.srem(key_, stuId);
            jedis.close();
        } catch (JedisDataException ex) {// multiple catch is not supported before java7
            ex.printStackTrace();
        } catch (NullPointerException ex) {
            ex.printStackTrace();
        } catch (JedisConnectionException ex) {
            ex.printStackTrace();
        }
    }

}


public class LogProcessBolt extends BaseBasicBolt {
    // private static Logger Log = LoggerFactory.getLogger(LogProcessBolt.class);
    private JedisAPI jedisAPI;
    private Matcher logMatcher;
    private Pattern p;
    private String targetKey;
    private String stuID, key_;
    private static final String logPattern = "^(\\[.*?\\])\\s+-\\s+(.*?)\\s+.*?-.*?\\s+\\w+_(.?)_(.*?)_(.*?)_(.*?).*@(.*?)\\s+\\[.*\\](\\s+.*)?";
    private String connectFlag;
    private String line;
    private static final int NThreads = 1000;
    private static final Executor exec = Executors.newCachedThreadPool();
    private JedisPool jedisPool;

    public void prepare(Map map, TopologyContext topologyContext) {

        p = Pattern.compile(logPattern);

        // 构造 reis 连接
        GenericObjectPoolConfig config = new GenericObjectPoolConfig();
        config.setMaxWaitMillis(MainConfig.MaxWaitMillis);
        config.setTestOnBorrow(MainConfig.TestOnBorrow);
        config.setTestOnReturn(MainConfig.TestOnReturn);
        config.setTestWhileIdle(MainConfig.TestWhileIdle);
        config.setMinEvictableIdleTimeMillis(MainConfig.MinEvictableIdleTimeMillis);
        config.setTimeBetweenEvictionRunsMillis(MainConfig.TimeBetweenEvictionRunsMillis);
        config.setNumTestsPerEvictionRun(MainConfig.NumTestsPerEvictionRun);
        config.setMaxTotal(MainConfig.MaxTotal);
        jedisPool = new JedisPool(config, MainConfig.redis_ip, MainConfig.redis_port, MainConfig.redis_timeout);
    }


    public void execute(Tuple tuple, BasicOutputCollector collector) {

        line = new String(tuple.getBinaryByField("logfile"));
        try {
            Matcher m = p.matcher(line);
            m.find();
            connectFlag = m.group(2);
            targetKey = m.group(3) + "_" + m.group(4);
            stuID = m.group(5);
            key_ = "LiveOnLineStuID_" + targetKey;
            /*
        * 接下来将 key 写入到 redis 中
        * */
            if (connectFlag.equals("Connect")) {
                // jedisAPI.insertToRedisBySet(key_, stuID);
                exec.execute(new InsertRedisThread(jedisPool.getResource(), key_, stuID));
            } else if (connectFlag.equals("Disconnect")) {
                //jedisAPI.removeFromRedisBySet(key_, stuID);
                exec.execute(new RemoveRedisThread(jedisPool.getResource(), key_, stuID));
            }
        } catch (Exception ex) {
        }


    }

    public void declareOutputFields(OutputFieldsDeclarer declarer) {
        declarer.declare(
                new Fields("OAOnlinePeopleCount"));
    }


    public void cleanup() {

    }

    public Map<String, Object> getComponentConfiguration() {
        return null;

    }
}

