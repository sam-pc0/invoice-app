import { templatesEnum, BIDProposal } from "@/type";

export default {
  generateInvoice(templateId) {
    let returnedObject = null;
    switch (Number(templateId)) {
      case templatesEnum.BID_PROPOSAL:
        returnedObject = new BIDProposal({
          id: 1,
          templateId: 1,
          name: "Test invoice",
          description: "A short description",
          number: 1,
          owner: {
            name: "",
            address: "",
            location: "",
            phone: "",
            altPhone: "2241",
            email: "",
            projectNameNAddress: "Test GT",
          },
          specificationNStimates:
            "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
          notIncluded:
            "Lorem ipsum dolor sit amet, consetetur sadipscing elitr,",
          totalSum: 500,
          withdrawnDays: 5,
          withdrawnDate: new Date(
            "Tue Jul 13 2021 00:00:00 GMT-0600 (Central Standard Time)"
          ),
        });
        break;

      case templatesEnum.INVOICE:
        returnedObject = new BIDProposal({
          id: 2,
          templateId: 2,
          name: "Test invoice",
          description: "A short description",
          number: 1,
          owner: {
            name: "",
            address: "",
            location: "",
            phone: "",
            altPhone: "2241",
            email: "",
            projectNameNAddress: "Test GT",
          },
          specificationNStimates:
            "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam",
          notIncluded:
            "Lorem ipsum dolor sit amet, consetetur sadipscing elitr,",
          totalSum: 500,
          withdrawnDays: 5,
          withdrawnDate: new Date(
            "Tue Jul 13 2021 00:00:00 GMT-0600 (Central Standard Time)"
          ),
        });
        break;
      case templatesEnum.CONTRACT_INVOICE:
        break;

      case templatesEnum.DAILY_N_MATERIAL_RECORD:
        break;
      default:
        console.info("no entro");
        break;
    }
    return returnedObject;
  },
};
