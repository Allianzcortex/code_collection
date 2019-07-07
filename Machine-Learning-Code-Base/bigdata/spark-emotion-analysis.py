from pyspark.sql import SparkSession
from pyspark import SQLContext,SparkConf

import numpy as np
import matplotlib
matplotlib.use('Agg')
import matplotlib.pyplot as plt

from textblob import TextBlob

# conf = SparkConf().setAppName("ReadCSV")
# sqlContext = SQLContext(conf=conf)
# sparkSession = SparkSession.builder.appName("spark-readfile").getOrCreate()
# df_load = sparkSession.read.csv('hdfs:///project/samples.csv')
# df_load.show()
def get_score(sentence):
    blob = TextBlob(sentence)
    return blob.sentences[0].sentiment.polarity

def run_spark():
    spark = SparkSession.builder.master("local").appName("Word Count").getOrCreate()
    df = spark.read.format("csv").option("header", "true").load("hdfs:///project/samples
    sqlContext=SQLContext(spark)
    sqlContext.registerDataFrameAsTable(df, "table1")
    sentences=sqlContext.sql("""SELECT `reviews.rating`,`reviews.text` FROM table1""").rdd
    sentences.collect()
    alist=sentences.map(lambda x:x[0] if x[0] is not None else 0).collect()
    blist=sentences.map(lambda x:get_score(x[1]) if x[1] is not None else 0).collect()
    return alist,blist

def draw_plots(alist,blist)
    alist,blist = run_spark()
    t_plt, = plt.plot(np.arange(1, len(alist)+1), alist, 'r')
    v_plt, = plt.plot(np.arange(1, len(alist)+1), blist)
    plt.title('NLP Emotion Analysis') #
    plt.xlabel('epoch') #
    plt.ylabel('score') #
    plt.legend((t_plt, v_plt), ('rating', 'score')) #
    plt.savefig("result.png")
    with open("x1.txt","w+") as input_:
        input_.write(','.join([str(x) for x in alist]))
        input_.write(','.join([str(x) for x in blist])) 

if __name__=='__main__':
    run_spark()