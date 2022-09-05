Feature: Delivery personnel registration

  As a system admin,
  want to  register delivery personnel
  so that they can start using the system to make deliveries.

  Scenario Outline: All required fields are given
    Given  system amdin is loged in
    When system admin sends "<first_name>" "<last_name>" "<phone_number>" "<vehicle_type>" "<vehicle_plate_number>" "<service_type>" "<email>"
    Then the result should be "<message>"
    Examples:
    | first_name   | last_name   | phone_number   | vehicle_type   | vehicle_plate_number   | service_type    | email                | message                        | 
    | Anwar        | Tuha        | +251909525760  | automobile     | 123456                 | parcel_delivery | anwartuha2@gmail.com | Driver registered successfully |
  
  Scenario Outline: Required fields are missing
   Given  system amdin is loged in
   When system admin sends "<first_name>" "<last_name>" "<phone_number>" "<vehicle_type>" "<vehicle_plate_number>" "<service_type>" "<email>"
   Then the result should be "<message>"
   Examples:
    | first_name   | last_name   | phone_number   | vehicle_type   | vehicle_plate_number   | service_type    | email                | message                                                    |
    |              | Tuha        | +251909525760  | automobile     | 123456                 | parcel_delivery | anwartuha2@gmail.com | Driver registration failed, first_name required            |
    | Anwar        |             | +251909525760  | automobile     | 123456                 | parcel_delivery | anwartuha2@gmail.com | Driver registration failed, last_name required             |
    | Anwar        | Tuha        |                | automobile     | 123456                 | parcel_delivery | anwartuha2@gmail.com | Driver registration failed, phone_number required          |
    | Anwar        | Tuha        | +251909525760  |                | 123456                 | parcel_delivery | anwartuha2@gmail.com | Driver registration failed, vehicle_type required          |
    | Anwar        | Tuha        | +251909525760  | automobile     |                        | parcel_delivery | anwartuha2@gmail.com | Driver registration failed, vehicle_plate_number required  |
    | Anwar        | Tuha        | +251909525760  | automobile     | 123456                 |                 | anwartuha2@gmail.com | Driver registration failed, vehicle_plate_number required  |

  
  Scenario Outline: Field inputs are invalid
    Given i have user 
      | first_name   | last_name   | phone_number   | vehicle_type   | vehicle_plate_number   | service_type    | email                |
      | Anwar        | Tuha        | +251909525760  | automobile     | 123456                 | parcel_delivery | anwartuha2@gmail.com |
    When system admin sents "<first_name>" "<last_name>" "<phone_number>" "<vehicle_type>" "<vehicle_plate_number>" "<service_type>" "<email>"
    Then the result should be "<message>"
    Examples:
    | first_name   | last_name   | phone_number   | vehicle_type   | vehicle_plate_number   | service_type    | email                | message                                                    |
    | Anwar        | Tuha        | +251           | automobile     | 123456                 | parcel_delivery | anwartuha2@gmail.com | Driver registration failed, invalid phone_number           |

  Scenario: default password
    Given i have user 
    | first_name   | last_name   | phone_number   | vehicle_type   | vehicle_plate_number   | service_type    | email                | message                        | 
    | Anwar        | Tuha        | +251909525760  | automobile     | 123456                 | parcel_delivery | anwartuha2@gmail.com | Driver registered successfully |
    
    When user sent "<phone_number>" and "<password>"
    Then result should be "<message>" 
    Examples: 
     |phone_number  | password |message  |
     |+251909525760 |AnwarTuha |success  |