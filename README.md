# Invoice app

### Códigios de Templates

```
1110 = Invoice
1100 = Bid Proposal
1000 = Contract Invoice
1001 = Daily Materials
```

## Bill Enpoint

- `/bills`
- `/bills/:id`

```json
// Json Reponse
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

