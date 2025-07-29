insert into users (name, surname, phone, company_id, department_id)
values (:name, :surname, :phone, :company_id, :department_id) returning id