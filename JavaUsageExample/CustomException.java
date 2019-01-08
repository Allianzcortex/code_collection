package payroll.utils;


public class CustomCheckedException extends Exception {

    /**
     * bwlow is available,but we will not know
     * the root cause of the exception
    public CustomException(String errorMessage){
        super(errorMessage);
    }
     **/

    /**
     below is the right way
     and the usage is like:
     try{

     } catch(GeneralException ex){
       if(isCustomReason(){
        throw new CustomCheckedException("Custom Reason ",ex);
     }
     }

     */
    public CustomCheckedException(String errorMessage,Throwable e){
        super(errorMessage,e);
    }
}

public class CustomUncheckedException extends RuntimeException