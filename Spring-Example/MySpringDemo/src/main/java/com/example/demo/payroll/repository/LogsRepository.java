package payroll.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import payroll.model.Logs;

import java.util.List;
interface
/**
 * 日志持久层
 */
public  LogsRepository extends JpaRepository<Long,Long> {

    /**
     * 查询最新的五条数据
     */
    @Query(value="select * from halo_logs order by log_created desc limit 5",nativeQuery = true)
    List<Logs> findTopFive();
}
