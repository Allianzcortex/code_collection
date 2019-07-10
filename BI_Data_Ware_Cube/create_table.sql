
crete database BIAssignment2;

use BIAssignment2


CREATE TABLE accident_fact(
	body_id int NOT NULL,
	class_id int NOT NULL,
	accident_id int NOT NULL,
	direction_id int NOT NULL,
	fuel_id int NOT NULL,
	vehicle_year_id int NOT NULL,
	number_of_occupants_id int NOT NULL,
	engine_cyclinders_id int NOT NULL,
	vehicle_make_id int NOT NULL
) 


create table direction_of_travel_4_dim (
	direction_id int identity(1,1) not null,
	direction_name varchar(50),
	primary key (direction_id)
)

create table fuel_type_5_dim (
	fuel_id int identity(1,1) not null,
	fuel_name varchar(50),
	primary key (fuel_id)
)


create table vehicle_year_6_dim (
	vehicle_year_id int identity(1,1) not null,
	vehicle_year varchar(50),
	primary key (vehicle_year_id)
)

create table number_of_occupants_7_dim (
	number_of_occupants_id int identity(1,1) not null,
	number_of_occupants int,
	primary key (number_of_occupants_id)
)

create table engine_cylinders_8_dim (
	engine_cylinders_id int identity(1,1) not null,
	engine_cylinder int,
	primary key (engine_cylinders_id)
)

create table vehicle_make_9_dim (
	vehicle_make_id int identity(1,1) not null,
	vehicle_make varchar(50),
	primary key(vehicle_make_id)
)