insert into repository_kind values (26, 'Radius recipes');
insert into repository_kind values (99, 'Python Lib');
insert into repository_kind values (98, 'Golang Lib');

---- create above / drop below ----

delete from repository_kind where repository_kind_id = 26;
delete from repository_kind where repository_kind_id = 99;
delete from repository_kind where repository_kind_id = 98;
