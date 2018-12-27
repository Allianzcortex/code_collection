package Test;

import io.netty.channel.ChannelHandlerContext;

import java.util.HashMap;
import java.util.Map;

/**
 * Created by hzcortex on 17-1-29.
 */
public class utils {
    static final Map<String, ChannelHandlerContext> container = new HashMap<String, ChannelHandlerContext>();
    static final Map<String, String> Connection = new HashMap<String, String>();

    public static void store(String addr, ChannelHandlerContext ctx) {
        container.put(addr, ctx);
    }
}
