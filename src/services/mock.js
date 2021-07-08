import { templatesEnum, BIDProposal } from "@/type";

export default {
  generateInvoice(templateId) {
    switch (Number(templateId)) {
      case templatesEnum.BID_PROPOSAL:
        return new BIDProposal({
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
          whitdrawnDays: 5,
          whitdrawnDate: "07-21-2020",
        });

      case templatesEnum.INVOICE:
        break;

      case templatesEnum.CONTRACT_INVOICE:
        break;

      case templatesEnum.DAILY_N_MATERIAL_RECORD:
        break;
      default:
        console.info("no entro");
        break;
    }
  },
};
