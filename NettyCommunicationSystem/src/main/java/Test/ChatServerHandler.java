package Test;

import io.netty.buffer.Unpooled;
import io.netty.channel.ChannelHandler.Sharable;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import io.netty.util.CharsetUtil;

import java.util.Collections;
import java.util.Map;

import static Test.utils.container;


/**
 * Handler implementation for echo server
 * use ConnectionMap to store
 * Created by hzcortex on 17-1-29.
 */

@Sharable
public class ChatServerHandler extends SimpleChannelInboundHandler<String> {
    @Override
    public void channelRead0(ChannelHandlerContext ctx, String msg) {
//        ByteBuf in = (ByteBuf) msg;
//        while (in.isReadable()) {
//            System.out.print((char) in.readByte());
//            System.out.flush();
//        }
        String addr = ctx.channel().remoteAddress().toString();
        System.out.println(addr);
        utils.store(addr, ctx);

        for (Map.Entry<String, ChannelHandlerContext> entry : container.entrySet()) {
            ctx.writeAndFlush(Unpooled.copiedBuffer(entry.getKey() + " \n", CharsetUtil.UTF_8).retain());
        }
        String content;
//        ByteBuf in = (ByteBuf) msg;
//        content = in.toString(StandardCharsets.UTF_8);
        content = msg;
        System.out.println("content== " + content);

        // 对传输来的数据如果为 send 那么和另一个 send 发送连接
        if (content.contains("send")) {
            String port = content.split(":")[1].trim();
            String key = "/127.0.0.1:" + port;
            ChannelHandlerContext newCtx = container.get(key);
            newCtx.writeAndFlush(content);
        }
    }

    @Override
    public void channelInactive(ChannelHandlerContext ctx) throws Exception {
        ctx.writeAndFlush("连接断开");
        container.values().remove(Collections.singleton(ctx));
        // 删除 container 和 connection 对应的链接
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        cause.printStackTrace();
        ctx.close();
    }
}
