package payroll.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import payroll.model.Logs;
import payroll.repository.LogsRepository;
import payroll.service.LogService;

import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import java.util.Date;
import java.util.List;

public class LogServiceImpl implements LogService{

    // 关于 Autowired 的作用，参考：
    // https://www.baeldung.com/spring-autowire
    @Autowired
    private LogsRepository logsRepository;

    @Override
    public void save(String logTitle, String logContent,
                     HttpServletRequest request){
        Logs logs=new Logs();
        logs.setLogTitle(logTitle);
        logs.setLogContent(logContent);
        logs.setLogRecord(new Date());
        logs.setLogIp();
        logsRepository.save(logs);
    }

    @Override
    public void removeAll(){
        logsRepository.deleteAll();
    }

    @Override
    public List<Logs> findTheLatestLogs(){
        return logsRepository.findTopFive();
    }

}
