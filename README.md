# Invoice app

### Códigios de Templates

```
1110 = Invoice
1100 = Bid Proposal
1000 = Contract Invoice
1001 = Daily Materials
```

## Bill Enpoint

- `/bills` | Post --> Retorna el Id de la Bill Insertada
- `/bills/:id` | Get
- `/bills?code=1100` | Put  --> En este caso BID | QueryParam
- `/bills?code=1110` | Put  --> En este caso Invoice | QueryParam

```json
// Json Reponse Get
{
  "id": 9,
  "name": "Test4",
  "description": "Testing",
  "owner": {
    "id": 0,
    "name": "",
    "location": "",
    "phone": "",
    "altphone": "",
    "project_name_address": "",
    "email": ""
  },
  "template": {
    "id": 1110,
    "template_code": 0
  },
  "last_edit": ""
}
```

```json
// Json Request BID Proposal Put
//Se espera que sea el ID de Bill (en este ejemplo la bill 9) y no un id generado para el BID
 {
  "id": 9, 
  "template_code": 1100,
	"name": "Test invoice",
	"description": "A short description",
	"number": 1,
    "owner": {
 				"name":"Test11",
				"location": "Jutiapa, Jutiapa 22001",
				"phone": "12345678",
				"altphone": "87654321",
				"project_name_address":"XD",
				"email":"correo@correo.com"
    },
     "specificationstimates":
     "Test11",
     "not_included":
     "Test11,",
     "totalsum": 500,
     "withdrawndays": 5,
     "withdrawndate": "Tue Jul 13 2021 00:00:00 GMT-0600 (Central Standard Time)"
 }
```

```json
// Json Request BID Proposal Put
//Se espera que sea el ID de Bill (en este ejemplo la bill 7) y no un id generado para el Invoice 
{
  "id": 7, 
  "template_code": 1110,
	"name": "Test invoice",
	"description": "A short description",
	"number": 1,
    "owner": {
 				"name":"Test13",
				"location": "Jutiapa, Jutiapa 22001",
				"phone": "12345678",
				"altphone": "87654321",
				"project_name_address":"xD",
				"email":"13@correo.com"
    },
     "item":[
			 {
				 "description":"Uno",
				 "amount":5
			 },
			 {
				 "description":"Dos",
				 "amount":5
			 }
		 ],
	 "total":10,
	 "date_submmitted":"Tue Jul 13 2021 00:00:00 GMT-0600 (Central Standard Time)"
 }
```

