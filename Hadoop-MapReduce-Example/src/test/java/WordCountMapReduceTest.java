/**
 * Created by allianzcortex on 17-4-23.
 */

import java.util.ArrayList;
import java.util.List;

import com.hzcortex.wordcount.WordCountMapper;
import com.hzcortex.wordcount.WordCountReducer;
import org.apache.hadoop.io.IntWritable;
import org.apache.hadoop.io.LongWritable;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mrunit.mapreduce.MapDriver;
import org.apache.hadoop.mrunit.mapreduce.MapReduceDriver;
import org.apache.hadoop.mrunit.mapreduce.ReduceDriver;
import org.junit.Before;
import org.junit.Test;


public class WordCountMapReduceTest {
    MapDriver<LongWritable, Text, Text, IntWritable> mapDriver;
    ReduceDriver<Text, IntWritable, Text, IntWritable> reduceDriver;
    MapReduceDriver<LongWritable, Text, Text, IntWritable, Text, IntWritable> mapReduceDriver;


    @Before
    public void setUp() {
        WordCountMapper mapper = new WordCountMapper();
        WordCountReducer reducer = new WordCountReducer();
        mapDriver=MapDriver.newMapDriver(mapper);
        reduceDriver=ReduceDriver.newReduceDriver(reducer);
        mapReduceDriver=MapReduceDriver.newMapReduceDriver(mapReduceDriver);
    }

    @Test
    public void testMapper() {
        mapDriver.withInput();
        mapDriver.withOutput();
        mapDriver.runTest();
    }

    @Test
    public void testReducer() {

    }

    @Test
    public void testMapReduce() {

    }


}
