create table tb_list_item (
  id varchar(64) not null constraint tb_list_item_pkey primary key,
  userName varchar(250) NOT NULL,
  description varchar(32) NOT NULL,
  value varchar(320) NOT NULL,  
  createdDate TIMESTAMP DEFAULT NOW()
);
