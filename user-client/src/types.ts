export type ChatMessage = {
  role: string;
  content: string;
};

export type ActionItemType = {
  title: string;
  details: string;
  type: string;
  completed: boolean;
  id: string;
};

type CompanyConfigValues = {
  companyName: string;
  industry: string;
  city: string;
  revenue: number;
  cogs: number;
  depreciation: number;
  longTermAssets: number;
  shortTermAssets: number;
  longTermLiability: number;
  shortTermLiability: number;
  operatingExpense: number;
  retainedEarnings: number;
  yearsInBusiness: number;
};
export type CompanyConfig = {
  riskValue: number;
  zValue: number;
  bankruptcyRisk: string;
  weightedBankruptcy: number;
  weightedBaseline: number;
  weightedRecs: number;
  active: CompanyConfigValues;
};

export type Reccomendation = {
  title: string;
  summary: string;
  details: string;
  type: string;
  weight: number;
  completed: boolean;
};
