<?php

$ffi = FFI::cdef(
    'extern int countTextLength(const char* text);',
    '/workspace/php-ffi-try'
);

$result = $ffi->countTextLength('hello');

echo($result);
echo "\n";
