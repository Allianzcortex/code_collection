package com.mcda.database.project.demo.utils;

import com.mcda.database.project.demo.model.Transaction;
import com.mcda.database.project.demo.repository.TransactionRepository;

import java.util.Calendar;
import java.util.Date;
import java.util.GregorianCalendar;
import java.util.List;

public class TransactionUtils {

    /**
     * Based on the formula : s = Sum*(1-2.5*DC/100)
     * define the X
     * the total purchase value of items by customer in the last five years
     * <p>
     * x>=500: Dc = 5
     * 400<=x<500: Dc = 4
     * 300<=x<400: Dc = 3
     * 200<=x<300: Dc = 2
     * 100<=x<200: Dc = 1
     * x<100: Dc = 0
     */

    public static int getDiscountCode(int customerId, TransactionRepository transactionRepository) {
        Date today = new Date();
        Calendar cal = new GregorianCalendar();
        cal.setTime(today);
        // 2. get transactions nearly 5 years
        cal.add(Calendar.YEAR, -5);
        Date yearago5 = cal.getTime();
        System.out.println(yearago5.getTime());

        List<Transaction> results = transactionRepository.findByCustomerIdAndTransactionDateAfter(
                customerId, yearago5);
        // TODO use java8 foreach to simplify the code
        int price_all = 0;
        int discount_code = 0;
        for (Transaction item : results) {
            price_all += item.getTotalPurchasePrice();
        }

        // Java switch doesn't support condition sentence
        // Certainly you can also extract the following login into utils function
        if (price_all < 100) {
            discount_code = 0;
        } else if (price_all >= 100 && price_all < 200) {
            discount_code = 1;
        } else if (price_all >= 200 && price_all < 300) {
            discount_code = 2;
        } else if (price_all >= 300 && price_all < 400) {
            discount_code = 3;
        } else if (price_all >= 400 && price_all < 500) {
            discount_code = 4;
        } else if (price_all >= 500) {
            discount_code = 5;
        }
        System.out.println("Currently discount_code is : " + discount_code);
        return discount_code;

    }

}
