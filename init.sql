create table tb_item (
  id BIGSERIAL not null constraint tb_list_item_pkey primary key,
  userName varchar(250) NOT NULL,
  description varchar(32) NOT NULL,
  value decimal NOT NULL,  
  createdDate TIMESTAMP DEFAULT NOW()
);
