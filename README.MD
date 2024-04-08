### Cadastrar Cliente

- **POST** `/customers`
- **Descrição**: Cadastra um novo cliente.
- **Corpo da Requisição**:
  ```json
  {
    "Name": "string",
    "Cnpj": "string"
  }
  ````

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "customer": {
        "CNPJ": "1234567890",
        "ID": 1,
        "Name": "Nome do cliente"
    }
  }
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Listar todos os Cliente

- **GET** `/customers`
- **Descrição**: Lista todos os cliente.

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "customers": [
        {
            "ID": 1,
            "CreatedAt": "2024-03-19T07:45:08.900425-03:00",
            "UpdatedAt": "2024-03-19T07:45:08.900425-03:00",
            "DeletedAt": null,
            "Name": "Nome do Cliente",
            "CNPJ": "1234567890",
            "Vehicles": []
        },
        // Outros clientes...
    ]
}
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Obter Cliente por  ID

- **GET** `/customers/:id`
- **Descrição**: Obtem cliente por id.

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "customer": {
        "ID": 1,
        "CreatedAt": "2024-03-19T07:45:08.900425-03:00",
        "UpdatedAt": "2024-03-19T07:45:08.900425-03:00",
        "DeletedAt": null,
        "Name": "Teste",
        "CNPJ": "1234567890",
        "Vehicles": []
    }
  }
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Atualizar Cliente

- **PUT** `/customers/:id`
- **Descrição**: Atualiza os detalhes de um cliente com base no ID.
- **Corpo da Requisição**:
  ```json
  {
    "Name": "string",
    "Cnpj": "string"
  }
  ````

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "customer": {
        "ID": 1,
        "CreatedAt": "2024-03-19T07:45:08.900425-03:00",
        "UpdatedAt": "2024-03-19T07:45:08.900425-03:00",
        "DeletedAt": null,
        "Name": "Teste",
        "CNPJ": "1234567890",
        "Vehicles": []
    }
  }
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Remover Cliente

- **DELETE** `/customers/:id`
- **Descrição**: Atualiza os detalhes de um cliente com base no ID.

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "customer": {
        "ID": 1,
        "CreatedAt": "2024-03-19T07:45:08.900425-03:00",
        "UpdatedAt": "2024-03-19T07:45:08.900425-03:00",
        "DeletedAt": "2024-03-19T07:45:08.900425-03:00",
        "Name": "Nome do Cliente",
        "CNPJ": "1234567890",
        "Vehicles": []
    }
  }
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Cadastrar Veículo

- **POST** `/vehicles`
- **Descrição**: Cadastra um novo veículo.
- **Corpo da Requisição**:
  ```json
  {
    "LicensePlate": "string",
    "Color": "string"
  }
  ````

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "vehicle": {
        "ID": 1,
        "CreatedAt": "2024-03-19T08:38:38.0651402-03:00",
        "UpdatedAt": "2024-03-19T08:38:38.0651402-03:00",
        "DeletedAt": null,
        "LicensePlate": "ABCD1234",
        "Color": "Color",
        "Customers": null
    }
}
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Listar todos os veículos

- **GET** `/vehicles`
- **Descrição**: Lista todos os veículos.

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "vehicles": [
      {
          "ID": 1,
          "CreatedAt": "2024-03-19T08:38:38.0651402-03:00",
          "UpdatedAt": "2024-03-19T08:38:38.0651402-03:00",
          "DeletedAt": null,
          "LicensePlate": "ABCD1234",
          "Color": "Color",
          "Customers": null
      },
      // outros veiculos...
  ]
}
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Obter Veiculo por  ID

- **GET** `/vehicles/:id`
- **Descrição**: Obtem veiculos por id.

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "vehicle": {
        "ID": 1,
        "CreatedAt": "2024-03-19T08:38:38.0651402-03:00",
        "UpdatedAt": "2024-03-19T08:38:38.0651402-03:00",
        "DeletedAt": null,
        "LicensePlate": "ABCD1234",
        "Color": "Color",
        "Customers": null
    }
  }
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Atualizar veiculos

- **PUT** `/vehicles/:id`
- **Descrição**: Atualiza os detalhes de um veiculo com base no ID.
- **Corpo da Requisição**:
  ```json
  {
    "LicensePlate": "string",
    "Color": "string"
  }
  ````

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "vehicle": {
        "ID": 1,
        "CreatedAt": "2024-03-19T08:38:38.0651402-03:00",
        "UpdatedAt": "2024-03-19T08:38:38.0651402-03:00",
        "DeletedAt": null,
        "LicensePlate": "ABCD1234",
        "Color": "Color",
        "Customers": null
    }
  }
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Remover Veiculo

- **DELETE** `/vehicles/:id`
- **Descrição**: Atualiza os detalhes de um veiculo com base no ID.

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "vehicle": {
        "ID": 1,
        "CreatedAt": "2024-03-19T08:38:38.0651402-03:00",
        "UpdatedAt": "2024-03-19T08:38:38.0651402-03:00",
        "DeletedAt": "2024-03-19T08:38:38.0651402-03:00",
        "LicensePlate": "ABCD1234",
        "Color": "Color",
        "Customers": null
    }
  }
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Adicionar veiculo para um  cliente 

- **POST** `/costumers/:id/vehicles/:vehicleID/links/:link`
- **Parametros opcionais**: `FORCE: true/false`
- **Descrição**: Adiciona um veiculo para um cliente com um determinado vinculo.

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  { 
    "customer": {
          "ID": 1,
          "CreatedAt": "2024-03-19T07:45:08.900425-03:00",
          "UpdatedAt": "2024-03-19T07:45:08.900425-03:00",
          "DeletedAt": "2024-03-19T07:45:08.900425-03:00",
          "Name": "Nome do Cliente",
          "CNPJ": "1234567890",
          "Vehicles": []
    },
    "vehicle": {
        "ID": 1,
        "CreatedAt": "2024-03-19T08:38:38.0651402-03:00",
        "UpdatedAt": "2024-03-19T08:38:38.0651402-03:00",
        "DeletedAt": "2024-03-19T08:38:38.0651402-03:00",
        "LicensePlate": "ABCD1234",
        "Color": "Color",
        "Customers": null
    }
  }
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Listar todos os veículos  de um cliente

- **GET** `/customers/:id/vehicles`
- **Descrição**: Lista todos os veículos  de um cliente.

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "vehicles": [
      {
          "ID": 1,
          "CreatedAt": "2024-03-19T08:38:38.0651402-03:00",
          "UpdatedAt": "2024-03-19T08:38:38.0651402-03:00",
          "DeletedAt": null,
          "LicensePlate": "ABCD1234",
          "Color": "Color",
          "Customers": null
      },
      // outros veiculos...
  ]
}
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Listar todos os veículos por viculo

- **GET** `/vehicles/links/:link`
- **Descrição**: Lista todos os veículos por vinculo.

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "vehicles": [
      {
          "ID": 1,
          "CreatedAt": "2024-03-19T08:38:38.0651402-03:00",
          "UpdatedAt": "2024-03-19T08:38:38.0651402-03:00",
          "DeletedAt": null,
          "LicensePlate": "ABCD1234",
          "Color": "Color",
          "Customers": null
      },
      // outros veiculos...
  ]
}
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Listar todos os veículos dado um cliente e um  vinculo

- **GET** `/customers/:id/vehicles/link/:link`
- **Descrição**: Lista todos os veículos por vinculo.

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "customers": [
        {
            "ID": 1,
            "CreatedAt": "2024-03-19T07:45:08.900425-03:00",
            "UpdatedAt": "2024-03-19T07:45:08.900425-03:00",
            "DeletedAt": null,
            "Name": "Nome do Cliente",
            "CNPJ": "1234567890",
            "Vehicles": []
        },
        // Outros clientes...
    ]
}
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Remover veículo de um cliente

- **DELETE** `/customers/:id/vehicles/:vehicleID`
- **Descrição**: Remove um veiculo de um cliente

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "customer": {
        "ID": 1,
        "CreatedAt": "2024-03-19T07:45:08.900425-03:00",
        "UpdatedAt": "2024-03-19T07:45:08.900425-03:00",
        "DeletedAt": null,
        "Name": "Teste",
        "CNPJ": "1234567890",
        "Vehicles": []
    }
  }
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```

### Remover todos od veículos de um cliente

- **DELETE** `/users/:id/vehicles`
- **Descrição**: Remove todos os veiculo de um cliente

## Respostas

### Sucesso

**Código:** 200

**Conteúdo:**

```json
  {
    "customer": {
        "ID": 1,
        "CreatedAt": "2024-03-19T07:45:08.900425-03:00",
        "UpdatedAt": "2024-03-19T07:45:08.900425-03:00",
        "DeletedAt": null,
        "Name": "Teste",
        "CNPJ": "1234567890",
        "Vehicles": []
    }
  }
```

### Erro

**Código:** 500

**Conteúdo:**

```json
  {
    "error": "mensagem de erro"
  }
```