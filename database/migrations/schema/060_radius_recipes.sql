insert into repository_kind values (26, 'Radius recipes');
insert into repository_kind values (27, 'Docker App');

---- create above / drop below ----

delete from repository_kind where repository_kind_id = 26;
delete from repository_kind where repository_kind_id = 27;
