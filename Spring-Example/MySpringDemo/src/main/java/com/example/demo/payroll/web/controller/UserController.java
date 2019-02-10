package payroll.web.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.validation.BindingResult;
import org.springframework.validation.ObjectError;
import org.springframework.web.bind.annotation.*;
import payroll.model.User;
import payroll.model.dto.HaloConst;
import payroll.model.dto.JsonResult;
import payroll.model.enums.ResultCodeEnum;
import payroll.service.UserService;

import javax.security.auth.login.Configuration;
import javax.servlet.http.HttpSession;
import javax.validation.Valid;

@Slf4j
@Controller
@RequestMapping(value = "/admin/profile")
public class UserController {

    @Autowired
    private UserService userService;

    /**
     * 获取用户信息并跳转
     *
     * @return 返回到前端的页面
     */
    @GetMapping
    public String profile() {
        return "admin/admin_profile";
    }


    @PostMapping(value = "save")
    // 根据官网说明，这个是
    //a method return value should be bound to the web
    //* response body. Supported for annotated handler methods.
    // 所以应该暂时用不上
    @ResponseBody
    public JsonResult saveProfile(@Valid @ModelAttribute User user,
                                  BindingResult result, HttpSession session){
        try{
            if(result.hasErrors()){
                for(ObjectError error:result.getAllErrors()){
                    return new JsonResult(ResultCodeEnum.FAIL.getCode(),
                            error.getDefaultMessage());
                }
            }

            userService.save(user);
            session.removeAttribute(HaloConst.USER_SESSION_KEY);
        } catch (Exception ex){
            log.error("", ex.getMessage());
            return new JsonResult(ResultCodeEnum.FAIL.getCode(),"edit fail");
        }
        return new JsonResult(ResultCodeEnum.SUCCESS.getCode(),"edit successful");
    }

}
