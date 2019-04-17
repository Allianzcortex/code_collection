package com.mcda.database.project.demo.repository;

import com.zaxxer.hikari.util.DriverDataSource;
import org.springframework.beans.factory.annotation.Value;


import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.jdbc.datasource.DriverManagerDataSource;
import org.springframework.jdbc.datasource.SimpleDriverDataSource;
import org.springframework.stereotype.Component;

import javax.transaction.Transactional;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.List;
import java.util.Map;

@Component
public class AllRepository {

    private JdbcTemplate jdbcTemplate;

    @Value("${spring.database.name}")
    private String databaseName;


    public AllRepository(@Value("${spring.datasource.url}") String databaseUrl
            , @Value("${spring.datasource.username}") String username,
                         @Value("${spring.datasource.password}") String password) {

        DriverManagerDataSource ds = new DriverManagerDataSource();
        ds.setDriverClassName("com.mysql.jdbc.Driver");
        System.out.println("url is " + databaseUrl);
        ds.setUrl(databaseUrl);
        ds.setUsername(username);
        ds.setPassword(password);
        this.jdbcTemplate = new JdbcTemplate(ds);
    }

    @Transactional
    public List<String> findAllAvailableTable() {
        // Just hard-coded the database name


        String query = "SELECT table_name FROM information_schema.tables WHERE table_schema ='" + databaseName + "';";
        return this.jdbcTemplate.query(query, new RowMapper<String>() {
            public String mapRow(ResultSet rs, int rowNum)
                    throws SQLException {
                return rs.getString(1);
            }
        });
    }

    @Transactional
    public List<Map<String, Object>> findAllTable(String tableName) {
        String query;
        if (tableName.equals("transactions")) {
            // show proper foramt date
            query = "select transaction_number,DATE_FORMAT(transaction_date, '%Y-%m-%d')," +
                    "total_purchase_price,customer_id from " + tableName;
        } else {
            query = "select * from " + tableName;
        }
        return this.jdbcTemplate.queryForList(query);
    }


    @Transactional
    public List<String> findAllTableFields(String tableName) {
        String query = "select COLUMN_NAME " +
                "  from information_schema.columns " +
                " where table_schema ='" + databaseName +
                "'   and table_name = '" + tableName + "'";

        query = "SHOW COLUMNS FROM " + tableName;

        return this.jdbcTemplate.query(query, new RowMapper<String>() {
            public String mapRow(ResultSet rs, int rowNum)
                    throws SQLException {
                return rs.getString(1);
            }
        });
    }
}
