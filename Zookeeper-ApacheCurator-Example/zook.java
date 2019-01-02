package Test;

import org.apache.zookeeper.*;

import java.io.IOException;
import java.util.List;
import java.util.concurrent.CountDownLatch;

/**
 * Created by hzcortex on 17-3-29.
 */


class ConnectionWatcher implements Watcher {
    private static final int SESSION_TIMEOUT = 5000;
    protected ZooKeeper zk;
    private CountDownLatch connectedSignal = new CountDownLatch(1);

    public void connect(String hosts) throws IOException, InterruptedException {
        zk = new ZooKeeper(hosts, SESSION_TIMEOUT, this);
        connectedSignal.await();
    }

    @Override
    public void process(WatchedEvent event) {
        if (event.getState() == Event.KeeperState.SyncConnected) {
            connectedSignal.countDown();
        }
    }

    public void close() throws InterruptedException {
        zk.close();
    }
}

class JoinGroup extends ConnectionWatcher {
    public void join(String groupName, String memberName) throws KeeperException,
            InterruptedException {
        String path = "/" + groupName + "/" + memberName;
        String createdPath = zk.create(path, null/*data*/, ZooDefs.Ids.OPEN_ACL_UNSAFE,
                CreateMode.PERSISTENT);
        System.out.println("Created " + createdPath);
    }
}

public class Zook implements Watcher {

    private static final int SESSION_TIMEOUT = 5000;
    private ZooKeeper zk;
    private CountDownLatch connectedSignal = new CountDownLatch(1);

    public void connect(String hosts) throws IOException, InterruptedException {
        zk = new ZooKeeper(hosts, SESSION_TIMEOUT, this);
        connectedSignal.await();
    }

    @Override
    public void process(WatchedEvent event) {
        if (event.getState() == Event.KeeperState.SyncConnected)
            connectedSignal.countDown();
    }

    public void create(String groupName) throws KeeperException, InterruptedException {
        String path = "/" + groupName;
        String createPath = zk.create(path, null, ZooDefs.Ids.OPEN_ACL_UNSAFE,
                CreateMode.PERSISTENT);
        System.out.println("创建 path " + path);
    }

    public void close() throws IOException, InterruptedException {
        zk.close();
    }

    /**
     * 打印出在 "/"+groupName 下所有的子节点列表
     *
     * @param groupName
     * @throws KeeperException
     * @throws InterruptedException
     */
    public void getList(String groupName) throws KeeperException, InterruptedException {
        String path = "/" + groupName;
        try {
            List<String> children = zk.getChildren(path, false);
            if (children.isEmpty()) {
                System.out.println("No Node Exiests");
                System.exit(1);
            }

            System.out.println("开始打印子节点");
            for (String child : children) {
                System.out.println(child);
            }
        } catch (KeeperException.NoNodeException ex) {
            System.err.println("Group not exists");
            System.exit(1);
        }
    }

    public void deletePath(String groupName) throws KeeperException, InterruptedException {
        String path = "/" + groupName;
        List<String> children = zk.getChildren(path, false);
        System.out.println("Begin to delete node");
        for (String child : children) {
            zk.delete(path + "/" + child, -1);
        }
        zk.delete(path, -1);
    }

    public static void main(String[] args) throws Exception {
        Zook zook = new Zook();
        zook.connect("10.97.8.11");
        zook.create("cortex");


        System.out.println("join: ");
        JoinGroup joinGroup = new JoinGroup();
        joinGroup.connect("10.97.8.11");
        joinGroup.join("cortex", "1");
        //Thread.sleep(Long.MAX_VALUE);

        zook.getList("cortex");
        zook.deletePath("cortex");
        System.out.println("reacquire: ");
        zook.getList("cortex");
        zook.close();
    }
}

