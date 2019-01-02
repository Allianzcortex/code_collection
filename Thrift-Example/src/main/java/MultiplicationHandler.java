/**
 * Created by allianzcortex on 17-4-19.
 */

import org.apache.thrift.TException;

public class MultiplicationHandler implements MultiplicationService.Iface {
    @Override
    public int multiply(int n1, int n2) throws TException {
        System.out.println("Multiply(" + n1 + "," + n2 + ")");
        return n1 * n2;
    }

}
