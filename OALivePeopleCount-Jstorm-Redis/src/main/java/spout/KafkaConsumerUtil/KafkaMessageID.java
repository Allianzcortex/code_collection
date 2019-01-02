package spout.KafkaConsumerUtil;

/**
 * Created by hzcortex on 16-11-15.
 */
public class KafkaMessageID {
    private int partition;
    private long offset;

    public KafkaMessageID(int partition, long offset) {
        this.setPartition(partition);
        this.setOffset(offset);
    }

    public int getPartition() {
        return partition;
    }

    private void setPartition(int partition) {
        this.partition = partition;
    }

    public long getOffset() {
        return offset;
    }

    private void setOffset(long offset) {
        this.offset = offset;
    }
}
