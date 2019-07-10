#INSERT INTO accident_fact(body_id,class_id,accident_id) VALUES (0,0,0) 

with open("insert_table.sql","w+") as input_:
    for index in range(1,1001):
        s="INSERT INTO accident_fact(body_id,class_id,accident_id,direction_id,fuel_id,vehicle_year_id,number_of_occupants_id,engine_cylinders_id,vehicle_make_id) VALUES ({0},{1},{2},{3},{4},{5},{6},{7},{8})\n".format(index,index,index,
            index,index,index,index,index,index)
        input_.write(s)

