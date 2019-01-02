package util;

/**
 * Created by hzcortex on 16-11-16.
 */

import org.apache.commons.pool2.impl.GenericObjectPoolConfig;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.exceptions.JedisConnectionException;
import redis.clients.jedis.exceptions.JedisDataException;

import config.MainConfig;

public class JedisAPI {
    private JedisPool jedisPool;
    private Jedis jedis = null;

    public JedisAPI() {
        preparePool();
        judgeAndGetJedis();
    }

    public void judgeAndGetJedis() {
        if (jedis == null) {
            jedis = jedisPool.getResource();
        }
    }

    public void preparePool() {
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

    public boolean insertToRedisBySet(String key, String value) {
        try {
            if (jedis == null) {
                judgeAndGetJedis();
            }
            jedis.select(MainConfig.redisDatabaseIndex);
            jedis.sadd(key, value);
            return true;
        } catch (JedisDataException ex) {// multiple catch is not supported before java7
            return false;
        } catch (NullPointerException ex) {
            ex.printStackTrace();
            return false;
        } catch (JedisConnectionException ex) {
            if (jedis != null) {
                jedis.close();
                jedis = null;
            }
            return false;
        }
    }

    public boolean removeFromRedisBySet(String key, String value) {
        try {
            if (jedis == null) {
                judgeAndGetJedis();
            }
            jedis.select(MainConfig.redisDatabaseIndex);
            jedis.srem(key, value);
            return true;
        } catch (JedisDataException ex) {// multiple catch is not supported before java7
            return false;
        } catch (NullPointerException ex) {
            ex.printStackTrace();
            return false;
        } catch (JedisConnectionException ex) {
            if (jedis != null) {
                jedis.close();
                jedis = null;
            }
            return false;
        }
    }
}
