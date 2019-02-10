package payroll.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import payroll.model.User;
import payroll.repository.UserRepository;
import payroll.service.UserService;

import java.util.Date;
import java.util.List;

// @Service 的作用是在于： A Spring bean in service layer should be annotated using @Service instead of @Component annotation
// @Controller, @Service, and @Repository are special types of @Component annotation.
@Service
public class UserServiceImpl implements UserService {
    @Autowired
    private UserRepository userRepository;

    @Override
    public void save(User user){
        userRepository.save(user);
    }

    // 理解了，这里用的是 Spring 自带的查询，通过 By 的方法把名字总结好了
    @Override
    public User userLoginByName(String userName, String userPass) {
        return userRepository.findByUserNameAndUserPassword(userName, userPass);
    }

    @Override
    public User userLoginByEmail(String userEmail, String userPass) {
        return userRepository.findByUserEmailAndUserPassword(userEmail, userPass);
    }

     @Override
    public User findByUserIdAndUserPass(Long userId, String userPass) {
        return userRepository.findByUserIdAndUserPassword(userId, userPass);
    }

    @Override
    public User findUser() {
        List<User> users=userRepository.findAll();
        if(users!=null && users.size()>0){
            return users.get(0);
        } else {
            return new User();
        }
    }

    @Override
    public void updateUserLoginEnable(String enable) {
        User user = this.findUser();
        user.setLoginEnable(enable);
        userRepository.save(user);
    }

    @Override
    public User updateUserLoginLastTime(Date lastDate) {
        User user = this.findUser();
        user.setLoginLast(lastDate);
        userRepository.save(user);
        return user;
    }

    @Override
    public Integer updateUserLoginErrorTimes() {
        User user = this.findUser();
        user.setLoginErrorTimes((user.getLoginErrorTimes() == null ? 0 : user.getLoginErrorTimes()) + 1);
        userRepository.save(user);
        return user.getLoginErrorTimes();
    }

    @Override
    public User updateUserNormal() {
        User user = this.findUser();
        user.setLoginEnable(TrueFalseEnum.TRUE.getDesc());
        user.setLoginErrorTimes(0);
        userRepository.save(user);
        return user;
    }


}
