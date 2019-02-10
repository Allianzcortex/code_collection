package payroll.service;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import payroll.model.Logs;

import javax.servlet.http.HttpServlet;
import java.util.List;

// service 这一层是实现所有业务逻辑的具体地方
// 默认是接口，在 impl 这里实现具体
public interface LogService {

    void save(String logTitle, String logContent,
              HttpServlet request);

    void removeAll();

    Page<Logs> findAll(Pageable pageable);

    List<Logs> findTheLatestLogs();
}
