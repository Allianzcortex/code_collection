package payroll.model;


import com.fasterxml.jackson.annotation.JsonIgnore;
import lombok.Data;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.validation.constraints.Email;
import javax.validation.constraints.NotBlank;
import java.io.Serializable;

@Data
@Entity
@Table(name = "halo_user")
public class User implements Serializable {
    private static final long serialVersionUID = -5144055068797033748L;

    @Id
    @GeneratedValue
    private Long userId;

    // 这个 validation 应该是用在表单验证
    // 自己 restful api 不用这样吧......
    @NotBlank(message = "用户名不能为空")
    // 在序列化的时候直接忽视这个值？
    @JsonIgnore
    private String userName;

    private String userDisplayName;

    @JsonIgnore
    private String userPassword;

    @Email(message = "邮箱格式不正确")
    private String userEmail;

    private String userAvatar;

    private String bio;

    @JsonIgnore
    private String loginEnable = "true";

    @JsonIgnore
    private Integer loginErrorTimes = 0;
}
