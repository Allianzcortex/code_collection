package com.mcda.database.project.demo.dto;

import lombok.*;

import java.util.List;
import java.util.Map;

@Setter
@Getter
@Data
@NoArgsConstructor
public class ReturnTables {
    private List<String> fields;
    private List<Map<String, Object>> values;
}
