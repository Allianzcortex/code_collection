package payroll.model.dto;

// 看起来 dto 就是来做这些事情的
public class JsonResult {

    /**
     * 返回的状态吗，0:失败,1:成功
     */
    private Integer code;

    private String message;

    private Object result;

    public JsonResult(Integer code){
        this.code=code;
    }

    public JsonResult(Integer code,String message){
        this.code=code;
        this.message=message;
    }

    public JsonResult(Integer code,String message,Object result){
        this.code=code;
        this.message=message;
        this.result=result;
    }

    public JsonResult(Integer code, Object result) {
        this.code = code;
        this.result = result;
    }
}
