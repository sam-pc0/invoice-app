export const templates = [
  {
    name: "BID Proposal",
    code: 1100,
  },
  {
    name: "Invoice",
    code: 1110,
  },

  {
    name: "Contract Invoice",
    code: 1000,
  },
  {
    name: "Daily Time & Material Record",
    code: 1001,
  },
];

export const templatesEnum = {
  BID_PROPOSAL: 1100,
  INVOICE: 1110,
  CONTRACT_INVOICE: 1000,
  DAILY_N_MATERIAL_RECORD: 1001,
};

export class Owner {
  constructor({
    id,
    name,
    address,
    location,
    phone,
    altPhone,
    projectNameNAddress,
    email,
  }) {
    this.id = id;
    this.name = name;
    this.address = address;
    this.location = location;
    this.phone = phone;
    this.altPhone = altPhone;
    this.projectNameNAddress = projectNameNAddress;
    this.email = email;
  }
}

export class BIDProposal {
  constructor({
    id,
    id_bid,
    templateCode,
    name,
    description,
    number,
    owner,
    specificationNStimates,
    notIncluded,
    totalSum,
    withdrawnDays,
    withdrawnDate,
  }) {
    this.id = id;
    this.id_bid = id_bid;
    this.templateCode = templateCode;
    this.name = name;
    this.description = description;
    this.number = number;
    this.owner = new Owner(owner);
    this.specificationNStimates = specificationNStimates;
    this.notIncluded = notIncluded;
    this.totalSum = totalSum;
    this.withdrawnDays = withdrawnDays;
    this.withdrawnDate = withdrawnDate;
  }
}

export class Item {
  constructor({ item, description, amount }) {
    this.item = item;
    this.description = description;
    this.amount = Number(amount);
  }
}

export class ItemList {
  constructor(itemsList) {
    let returnedList = [];
    itemsList.forEach((item) => {
      returnedList.push(new Item(item));
    });
    return returnedList;
  }
}

export class Invoice {
  constructor({
    id,
    template_code,
    last_edit,
    name,
    description,
    owner,
    items,
    total,
    dateSubmitted,
  }) {
    this.id = id;
    this.last_edit = new Date(last_edit);
    this.template_code = template_code;
    this.name = name;
    this.description = description;
    this.owner = new Owner(owner);
    this.items = new ItemList(items);
    this.total = total;
    this.dateSubmitted =
      dateSubmitted !== "" ? new Date(dateSubmitted) : new Date();
  }
}

export class ContractInvoice {
  constructor({
    id,
    templateCode,
    number,
    owner,
    itemsList,
    total,
    dateSubmitted,
    originalContractAmount,
  }) {
    this.id = id;
    this.templateCode = templateCode;
    this.number = number;
    this.owner = owner;
    this.itemsList = itemsList;
    this.total = total;
    this.dateSubmitted = dateSubmitted;
    this.originalContractAmount = originalContractAmount;
  }
}

export class Materials {
  constructor({ number, qty, description, atSign, amount }) {
    this.number = number;
    this.qty = qty;
    this.description = description;
    this.atSign = atSign;
    this.amount = amount;
  }
}

export class Workers {
  constructor(worker, hours, rate, amount) {
    this.worker = worker;
    this.hours = hours;
    this.rate = rate;
    this.amount = amount;
  }
}

export class MaterialsNWorkersItem {
  constructor(material, worker) {
    this.material = material;
    this.worker = worker;
  }
}

export class MaterialsNWorkers {
  constructor({
    id,
    templateCode,
    number,
    owner,
    materialNworkersList,
    totalMaterials,
    totalLabor,
    descriptionOfWork,
    otherDescription,
    dateSubmitted,
  }) {
    this.id = id;
    this.templateCode = templateCode;
    this.number = number;
    this.owner = owner;
    this.materialNworkersList = materialNworkersList;
    this.totalMaterials = totalMaterials;
    this.totalLabor = totalLabor;
    this.descriptionOfWork = descriptionOfWork;
    this.otherDescription = otherDescription;
    this.dateSubmitted = dateSubmitted;
  }
}
