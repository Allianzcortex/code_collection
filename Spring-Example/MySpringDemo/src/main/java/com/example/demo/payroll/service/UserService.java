package payroll.service;

import payroll.model.User;

import java.util.Date;

public interface UserService {

    // 保存用户
    void save(User user);

    void userLoginByName(String userName,String userPassword);

    void userLoginByEmail(String userEmail,String userPassword);

    // 查找所有的用户
    User findAllUser();

    User findUserByUserIdAndUserPass(Long userId,String userPass);

    // 修改禁用状态
    void updatUserLoginEnable(String enable);

    // 修改最后登录时间
    User updateUserLoginLastTime(Date lastDate);

    Integer updateUserLoginErrorTimes();
}
