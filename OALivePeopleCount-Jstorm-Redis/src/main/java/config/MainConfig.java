package config;

/**
 * Created by hzcortex on 16-11-16.
 */
public class MainConfig {
    //public static String redis_ip = "redis-oa-live-people-count";
    public static String redis_ip = "127.0.0.1";
    public static int redis_port = 6379;
    public static int redis_timeout = 400000;
    public static int MaxWaitMillis = 1000;
    public static Boolean TestOnBorrow = false;
    public static Boolean TestOnReturn = false;
    public static Boolean TestWhileIdle = true;
    public static int MinEvictableIdleTimeMillis = 120000;
    public static int TimeBetweenEvictionRunsMillis = 60000;
    public static int NumTestsPerEvictionRun = -1;
    public static int MaxTotal = 1000;
    public static int redisDatabaseIndex = 1;
}
