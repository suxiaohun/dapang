package rsa_tools

func TestPrivateKeyPem() string {
	pem := `
-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEA0KU43pDibNlu8VCkzKR7pLEkpZjnk0cXx9ZCeBBQz56HXZLo
27gc7PBplzv7M3KKU8bsFW3bKe5fMEJ4KAmMlU+zgrIwdpi2jMOUS6DMCiNi/naJ
LzNxrew6LzLpHp4M4v6qUrv9nlW2Wvdgu89SKJlZCdHakGYqP65bSE/JS35IHNEZ
31QQ/JpVmvK8QYY3GGlURK6D2C5g74ijytEWtdKjr6p9feDeBXbaMrrb+9p/kekA
piDmZHMtDuGMiLLeQcKddl2ynO/z39sYDNJM7Iw/5eFvnYvCwr15NzQHFQe5yWNl
7l0d4O/xS2x7Hrerlktr/1urHCCMSs4oLyqAYwIDAQABAoIBAQDNjdlHPEo2nsh0
KdUSYAIIMRU/qDKNoH7HPElj6/8RbMdD4xZSUUL+E6STgF5tQy0QeGvieMZwuw1A
iFvEZhQAjfL2kgK49J70c0LId3RNkKtYswNYY0Bd7tsr4OZWabOcdvTqaZblxN/R
IuJEZW413U4hJIRFusFo16fKyAAyxEFM+0Q1CM7EJdA4S9EIW3K8NcDp6og/pBlW
I0h+GwddqyfOO8kXCi43uze+d8Tj4Qmvnv+Z/1uxSI5ZNlVObK2GYNoPJ550eZ57
UlcEjZEih5kn7Jl7A+3VmceF7d/Th7jOxCS8FH3Sl/dx7Py+6mxPHO6GyQX3+MfX
+MatwlrhAoGBAODbpPK/0kDXv3T2AnhZX06lCRjEvNrJ4YcnCvC5kM7h5wO62zmD
bk74MBmKb+X48U3yaBNJ6AOSkWvATNQctKXEdH8b8UGiFUClyg7KNkUY0BEgTbkB
BXgkgVdXfMWIjSBhqHYuEepz445aCG1jFyMhvppRwfg6fXurEHqdZPWLAoGBAO2K
wmlF2hbTuOF17BS5MoNUkQEtSQPiR7VSGHjhY2dxAlJ732/uKLoKP9kE6mbIkC7K
G6x7Zg5H6r/A7U2MSJZRxs84lIdDMLYsSgLNPr6sdJ3rMoDfSdrE5whNHFzkr0Si
COeteq9x4OYgkFzz5I77v9ts0KxuMuncG8uIxmuJAoGBAMTSvmHWUp57u5UPuj3K
TUiz4XCS1ZTV3dnF7Sel5wGQb/yt1TNUvnvVd5UDvssYy3AlXcB0kuh9pVkiLJ4G
uoHXovmYZefDjsLV3DZx/mshOw+irSXhpP3zRlxDsm3eB1pu4V0BGqGU630xherq
L9iKsyK1qgo0eJ6FsJkDqNndAoGBAOauRjYtm6lnWre46NcHMsLPMTJPZATdjn3L
h/B8ESi3WmX33fi489wXcAG1AF1okHyf0VgjTey4coRdOxaAEU3JhKRgcf0Zhwii
4gzufOx2VZm2eokkTnihlys8lzqu6nZgVP2IPwh6MsnI/ekHAQLKAHt6pPuO2Oq6
0lNKNEbhAoGAYswEp7IGH1vONTLo/WTuJbwn2K8LDNJkodOMsrsqUVgcWKhE7SS7
kiksvkMXvYjCM9GBRPi52Odse/clUnpOOQYNz5quZuyc15b/F+E4s6YA/QtWz36g
yPx4f2biHKcpb9rG+LcksXNhfwxfwvstv720GvcpLxAEanT2x78bgI8=
-----END RSA PRIVATE KEY-----
`
	return pem
}

func TestPublicKeyPem() string  {
	pem := `
-----BEGIN RSA PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0KU43pDibNlu8VCkzKR7
pLEkpZjnk0cXx9ZCeBBQz56HXZLo27gc7PBplzv7M3KKU8bsFW3bKe5fMEJ4KAmM
lU+zgrIwdpi2jMOUS6DMCiNi/naJLzNxrew6LzLpHp4M4v6qUrv9nlW2Wvdgu89S
KJlZCdHakGYqP65bSE/JS35IHNEZ31QQ/JpVmvK8QYY3GGlURK6D2C5g74ijytEW
tdKjr6p9feDeBXbaMrrb+9p/kekApiDmZHMtDuGMiLLeQcKddl2ynO/z39sYDNJM
7Iw/5eFvnYvCwr15NzQHFQe5yWNl7l0d4O/xS2x7Hrerlktr/1urHCCMSs4oLyqA
YwIDAQAB
-----END RSA PUBLIC KEY-----
`
	return pem
}


func TestPrivateKeyStr() string {
	pem := `MIIEpQIBAAKCAQEA0KU43pDibNlu8VCkzKR7pLEkpZjnk0cXx9ZCeBBQz56HXZLo27gc7PBplzv7M3KKU8bsFW3bKe5fMEJ4KAmMlU+zgrIwdpi2jMOUS6DMCiNi/naJLzNxrew6LzLpHp4M4v6qUrv9nlW2Wvdgu89SKJlZCdHakGYqP65bSE/JS35IHNEZ31QQ/JpVmvK8QYY3GGlURK6D2C5g74ijytEWtdKjr6p9feDeBXbaMrrb+9p/kekApiDmZHMtDuGMiLLeQcKddl2ynO/z39sYDNJM7Iw/5eFvnYvCwr15NzQHFQe5yWNl7l0d4O/xS2x7Hrerlktr/1urHCCMSs4oLyqAYwIDAQABAoIBAQDNjdlHPEo2nsh0KdUSYAIIMRU/qDKNoH7HPElj6/8RbMdD4xZSUUL+E6STgF5tQy0QeGvieMZwuw1AiFvEZhQAjfL2kgK49J70c0LId3RNkKtYswNYY0Bd7tsr4OZWabOcdvTqaZblxN/RIuJEZW413U4hJIRFusFo16fKyAAyxEFM+0Q1CM7EJdA4S9EIW3K8NcDp6og/pBlWI0h+GwddqyfOO8kXCi43uze+d8Tj4Qmvnv+Z/1uxSI5ZNlVObK2GYNoPJ550eZ57UlcEjZEih5kn7Jl7A+3VmceF7d/Th7jOxCS8FH3Sl/dx7Py+6mxPHO6GyQX3+MfX+MatwlrhAoGBAODbpPK/0kDXv3T2AnhZX06lCRjEvNrJ4YcnCvC5kM7h5wO62zmDbk74MBmKb+X48U3yaBNJ6AOSkWvATNQctKXEdH8b8UGiFUClyg7KNkUY0BEgTbkBBXgkgVdXfMWIjSBhqHYuEepz445aCG1jFyMhvppRwfg6fXurEHqdZPWLAoGBAO2KwmlF2hbTuOF17BS5MoNUkQEtSQPiR7VSGHjhY2dxAlJ732/uKLoKP9kE6mbIkC7KG6x7Zg5H6r/A7U2MSJZRxs84lIdDMLYsSgLNPr6sdJ3rMoDfSdrE5whNHFzkr0SiCOeteq9x4OYgkFzz5I77v9ts0KxuMuncG8uIxmuJAoGBAMTSvmHWUp57u5UPuj3KTUiz4XCS1ZTV3dnF7Sel5wGQb/yt1TNUvnvVd5UDvssYy3AlXcB0kuh9pVkiLJ4GuoHXovmYZefDjsLV3DZx/mshOw+irSXhpP3zRlxDsm3eB1pu4V0BGqGU630xherqL9iKsyK1qgo0eJ6FsJkDqNndAoGBAOauRjYtm6lnWre46NcHMsLPMTJPZATdjn3Lh/B8ESi3WmX33fi489wXcAG1AF1okHyf0VgjTey4coRdOxaAEU3JhKRgcf0Zhwii4gzufOx2VZm2eokkTnihlys8lzqu6nZgVP2IPwh6MsnI/ekHAQLKAHt6pPuO2Oq60lNKNEbhAoGAYswEp7IGH1vONTLo/WTuJbwn2K8LDNJkodOMsrsqUVgcWKhE7SS7kiksvkMXvYjCM9GBRPi52Odse/clUnpOOQYNz5quZuyc15b/F+E4s6YA/QtWz36gyPx4f2biHKcpb9rG+LcksXNhfwxfwvstv720GvcpLxAEanT2x78bgI8=`
	return pem
}

func TestPublicKeyStr() string  {
	pem := `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0KU43pDibNlu8VCkzKR7pLEkpZjnk0cXx9ZCeBBQz56HXZLo27gc7PBplzv7M3KKU8bsFW3bKe5fMEJ4KAmMlU+zgrIwdpi2jMOUS6DMCiNi/naJLzNxrew6LzLpHp4M4v6qUrv9nlW2Wvdgu89SKJlZCdHakGYqP65bSE/JS35IHNEZ31QQ/JpVmvK8QYY3GGlURK6D2C5g74ijytEWtdKjr6p9feDeBXbaMrrb+9p/kekApiDmZHMtDuGMiLLeQcKddl2ynO/z39sYDNJM7Iw/5eFvnYvCwr15NzQHFQe5yWNl7l0d4O/xS2x7Hrerlktr/1urHCCMSs4oLyqAYwIDAQAB`
	return pem
}
