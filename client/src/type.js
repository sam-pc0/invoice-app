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

export class TotalsData {
  constructor({ subTotal, total, taxRate, tax }) {
    this.subTotal = subTotal;
    this.total = total;
    this.taxRate = taxRate;
    this.tax = tax;
  }
}

export class BIDProposal {
  constructor({
    id,
    template_code,
    name,
    description,
    owner,
    specificationNStimates,
    notIncluded,
    subTotal,
    total,
    taxRate,
    tax,
    submittedBy,
    withdrawnDays,
    withdrawnDate,
  }) {
    this.id = id;
    this.template_code = template_code;
    this.name = name;
    this.description = description;
    this.owner = new Owner(owner);
    this.specificationNStimates = specificationNStimates;
    this.notIncluded = notIncluded;
    this.withdrawnDays = withdrawnDays;
    this.subTotal = subTotal;
    this.total = total;
    this.taxRate = taxRate;
    this.tax = tax;
    this.submittedBy = submittedBy;
    this.withdrawnDate =
      withdrawnDate !== "" ? new Date(withdrawnDate) : new Date();
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
    subTotal,
    total,
    taxRate,
    tax,
    dateSubmitted,
  }) {
    this.id = id;
    this.last_edit = new Date(last_edit);
    this.template_code = template_code;
    this.name = name;
    this.description = description;
    this.owner = new Owner(owner);
    this.items = new ItemList(items);
    this.subTotal = subTotal;
    this.total = total;
    this.taxRate = taxRate;
    this.tax = tax;
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
