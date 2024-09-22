insert into repository_kind values (26, 'Radius recipes');
insert into repository_kind values (27, 'Golang Lib');
insert into repository_kind values (28, 'Python Lib');

---- create above / drop below ----

delete from repository_kind where repository_kind_id = 26;
delete from repository_kind where repository_kind_id = 27;
delete from repository_kind where repository_kind_id = 28;
