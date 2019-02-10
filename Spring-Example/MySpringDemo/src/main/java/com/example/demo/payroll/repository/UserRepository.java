package payroll.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import payroll.model.User;

public interface UserRepository extends JpaRepository<User, Long> {

    // 根据用户名和密码查询
    // 这里应该是等价于
    // @query("select user from user_halo where userName=userName
    // and userPassword=userPassword,query=trye))
    User findByUserNameAndUserPassword(String userName,
                                       String userPassword);

    User findByUserEmailAndUserPassword(String userEmail,
                                        String userPassword);

    User findByUserIdAndUserPassword(Long usreId,
                                     String userPassword);
}
