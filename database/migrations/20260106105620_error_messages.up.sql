SET
    FOREIGN_KEY_CHECKS = 0;

INSERT INTO
  error_messages (code, component, response_type, one, other)
VALUES
  (
    'CNF0000000000',
    'application',
    'InternalServerError',
    'Error application config failed',
    'Error application config failed'
  ),
  (
    'CNF0000000010',
    'application',
    'InternalServerError',
    'Error in internal server',
    'Error in internal server'
  ),
  (
    'CNF0000000020',
    'repository',
    'InternalServerError',
    'Error database operation',
    'Error database operation'
  ),
  (
    'CNF0000000070',
    'application',
    'InternalServerError',
    'Failed to parse public key.',
    'Failed to parse public key.'
  ),
  (
    'CNF0000000080',
    'application',
    'InternalServerError',
    'Failed to parse private key.',
    'Failed to parse private key.'
  ),
  (
    'REQ0000000000',
    'controller',
    'BadRequest',
    'Error faild to extract data from request.',
    'Error faild to extract data from request.'
  ),
  (
    'REQ0000000020',
    'controller',
    'InternalServerError',
    'Error fetching user from gin context with name {{.name}}',
    'Error fetching user from gin context with name {{.name}}'
  ),
  (
    'REQ0000000040',
    'controller',
    'Unauthorized',
    'Error in fetching authtoken from request header.',
    'Error in fetching authtoken from request header.'
  ),
  (
    'REQ0000000050',
    'controller',
    'Unauthorized',
    'Error authtoken is invalid in the request header.',
    'Error authtoken is invalid in the request header.'
  ),
  (
    'REQ0000000060',
    'controller',
    'BadRequest',
    'Error request body validation failed.',
    'Error request body validation failed.'
  ),
  (
    'REQ0000000070',
    'service',
    'InternalServerError',
    'Error generating post request.',
    'Error generating post request.'
  ),
  (
    'REQ0000000090',
    'service',
    'InternalServerError',
    'Error executing request.',
    'Error executing request.'
  ),
  (
    'REQ0000000100',
    'service',
    'InternalServerError',
    'Error reading response body.',
    'Error reading response body.'
  ),
  (
    'GTK0000000000',
    'controller',
    'BadRequest',
    'Cannot generate token.',
    'Cannot generate token.'
  ),
  (
    'OTP0000000000',
    'repository',
    'NotFound',
    'No otp record found for application Number {{.appNumber}}.',
    'No otp record found for application Number {{.appNumber}}.'
  ),
  (
    'OTP0000000030',
    'service',
    'BadRequest',
    'Error invalid otp type {{.otpType}}',
    'Error invalid otp type {{.otpType}}'
  ),
  (
    'OTP0000000040',
    'service',
    'BadRequest',
    'Application is not in state to allow this operation.',
    'Application is not in state to allow this operation.'
  ),
  (
    'OTP0000000050',
    'repository',
    'BadRequest',
    'Error otp is empty',
    'Error otp is empty'
  ),
  (
    'REPO0000000000',
    'repository',
    'NotFound',
    'No record for search paramters {{.params}}.',
    'No record for search paramters {{.params}}.'
  ),
  (
    'USR0000000000',
    'repository',
    'NotFound',
    'Error user with {{.user_id}} not found.',
    'Error user with {{.user_id}} not found.'
  ),
  (
    'USR0000000010',
    'service',
    'AlreadyExists',
    'Error user with {{.email}} already exists.',
    'Error user with {{.email}} already exists.'
  ),
  (
    'USR0000000030',
    'service',
    'NotFound',
    'User permsssion not found {{.code}}.',
    'User permsssion not found {{.code}}.'
  ),
  (
    'USR0000000040',
    'controller',
    'BadRequest',
    'User role not created for userId {{.userId}}.',
    'User role not created for userId {{.userId}}.'
  ),
  (
    'USR0000000050',
    'service',
    'BadRequest',
    'Invalid password.',
    'Invalid password.'
  ),
  (
    'ROLE000000000',
    'repository',
    'NotFound',
    'Error role id {{.role_id}} not found.',
    'Error role id {{.role_id}} not found..'
  ),
  (
    'FELIX00000000',
    'service',
    'BadRequest',
    'Error Sending Notification.',
    'Error Sending Notification.'
  ),
  (
    'COMMON0000000',
    'service',
    'BadRequest',
    'Error Marshal Data.',
    'Error Marshal Data.'
  ),
  (
    'CRYPTO0000000',
    'service',
    'InternalServerError',
    'Error while encryption.',
    'Error while encryption.'
  ),
  (
    'CRYPTO0000010',
    'service',
    'InternalServerError',
    'Error while decryption.',
    'Error while decryption.'
  ),
  (
    'PMS0000000000',
    'repository',
    'NotFound',
    'Error permission id {{.permission_id}} not found.',
    'Error permission id {{.permission_id}} not found..'
  ),
  (
    'RP00000000000',
    'repository',
    'NotFound',
    'Error role permission not found for permission id {{.permission_id}}.',
    'Error role permission not found for permission id {{.permission_id}}.'
  ),
  (
    'RRQ0000000000',
    'repository',
    'BadRequest',
    'Error record request not found.',
    'Error record request not found.'
  );
  
SET
    FOREIGN_KEY_CHECKS = 1;
