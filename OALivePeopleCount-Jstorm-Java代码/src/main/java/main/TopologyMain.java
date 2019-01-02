package main;

/**
 * Created by hzcortex on 16-11-15.
 */

import java.util.HashMap;
import java.util.Map;

import backtype.storm.Config;
import backtype.storm.LocalCluster;
import backtype.storm.StormSubmitter;
import backtype.storm.generated.AlreadyAliveException;
import backtype.storm.generated.InvalidTopologyException;
import backtype.storm.topology.TopologyBuilder;

import spout.KafkaSpout;
import spout.SendLogSpout;
import bolt.LogProcessBolt;


public class TopologyMain {
    public static void main(String[] args) throws AlreadyAliveException, InvalidTopologyException {


        TopologyBuilder builder = new TopologyBuilder();


        KafkaSpout SendSpout = SendLogSpout.createSpout();

        builder.setSpout("SendSpout", SendSpout, 2);

        builder.setBolt("LogProcessBolt", new LogProcessBolt(), 3).shuffleGrouping("SendSpout");

        Config pconf = new Config();
        pconf.setNumAckers(0);
        pconf.setDebug(true);
//        pconf.setNumWorkers(3);
//        String topologyName = "OALivePeopleCount";
//        StormSubmitter.submitTopology(topologyName, pconf,
//                builder.createTopology());
        LocalCluster cluster = new LocalCluster();
        cluster.submitTopology("userlogCaltest", pconf,
                builder.createTopology());
//        pconf.setNumWorkers(3);
//        StormSubmitter.submitTopology(args[0], pconf,
//                builder.createTopology());
        /*if (args != null && args.length > 0) {
            pconf.setNumWorkers(3);
            StormSubmitter.submitTopology(args[0], conf,
                    builder.createTopology());

        } else {
            LocalCluster cluster = new LocalCluster();
            cluster.submitTopology("userlogCaltest", conf,
                    builder.createTopology());
            //Utils.sleep(10);
            //cluster.killTopology("userlogCaltest");
            //cluster.shutdown();
        }*/

    }
}


