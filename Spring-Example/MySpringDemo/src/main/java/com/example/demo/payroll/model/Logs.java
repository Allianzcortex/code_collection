package payroll.model;

import lombok.Data;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;
import java.io.Serializable;
import java.util.Date;

/**
 * <pre>
 *
 * </pre>
 */

@Data
@Entity
@Table(name="halo_logs")
public class Logs implements Serializable{

    @Id
    @GeneratedValue
    private Long logId;

    /**
     * log content
     */

    private String logTitle;

    private String logContent;

    private String logIp;

    private Date logRecord;


}
