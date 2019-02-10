package payroll.model;

import lombok.Data;

import javax.persistence.Entity;
import javax.persistence.ManyToMany;
import javax.persistence.Table;

@Data
@Entity
//@ManyToMany() // 如何限制对应关系
//@Table() // 不知道这个作用是哪里
public class Job {
}
