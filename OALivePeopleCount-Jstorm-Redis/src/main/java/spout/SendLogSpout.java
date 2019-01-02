package spout;

import java.util.HashMap;
import java.util.Map;

import config.KafkaSpoutConfig;
/**
 * Created by hzcortex on 16-11-15.
 */
public class SendLogSpout {
    public static KafkaSpout createSpout() {
        Map<String,String> Maconf = new HashMap<String, String>();
        Maconf.put("kafka.zookeeper.hosts","10.10.8.95:2181,10.10.8.96:2181,10.10.8.97:2181");
        Maconf.put("kafka.topic", "new_live_student_log");
        Maconf.put("kafka.broker.hosts", "10.10.8.101,10.10.8.102,10.10.8.103");
        Maconf.put("kafka.client.id", "new_live_student_log"); // 使用 kafka_client_id 会造成无法 fetch message
        Maconf.put("kafka.broker.partitions", "1");
        Maconf.put("metadata.broker.list", "10.10.8.101:9092,10.10.8.102:9092,10.10.8.103:9092");

        KafkaSpoutConfig MakafkaConfig = new KafkaSpoutConfig();
        MakafkaConfig.configure(Maconf);

        return new KafkaSpout(MakafkaConfig);
    }
}
