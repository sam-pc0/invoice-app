export const templates = [
  {
    name: "BID Proposal",
    id: 1,
  },
  {
    name: "Invoice",
    id: 2,
  },

  {
    name: "Contract Invoice",
    id: 3,
  },
  {
    name: "Daily Time & Material Record",
    id: 4,
  },
];

export const templatesEnum = {
  BID_PROPOSAL: 1,
  INVOICE: 2,
  CONTRACT_INVOICE: 3,
  DAILY_N_MATERIAL_RECORD: 4,
};

export class Owner {
  constructor({
    name,
    address,
    location,
    phone,
    altPhone,
    projectNameNAddress,
    email,
  }) {
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
    templateId,
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
    this.templateId = templateId;
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
  constructor(item, description, amount) {
    this.item = item;
    this.description = description;
    this.amount = amount;
  }
}

export class Invoice {
  constructor({
    id,
    templateId,
    number,
    owner,
    itemsList,
    total,
    dateSubmitted,
  }) {
    this.id = id;
    this.templateId = templateId;
    this.number = number;
    this.owner = owner;
    this.itemsList = itemsList;
    this.total = total;
    this.dateSubmitted = dateSubmitted;
  }
}

export class ContractInvoice {
  constructor({
    id,
    templateId,
    number,
    owner,
    itemsList,
    total,
    dateSubmitted,
    originalContractAmount,
  }) {
    this.id = id;
    this.templateId = templateId;
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
    templateId,
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
    this.templateId = templateId;
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
