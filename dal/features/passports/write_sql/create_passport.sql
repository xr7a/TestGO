insert into passports (type, number)
values (:type, :number) returning id, type, number