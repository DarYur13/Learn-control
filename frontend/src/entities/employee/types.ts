export interface EmployeeBaseInfo {
    id: number;
    fullname: string;
    birthdate: string;
  }
  
  export interface PersonalCard {
    employeeID: number;
    fullname: string;
    birthdate: string;
    department: string;
    position: string;
    employmentDate: string;
    snils: string;
    trainings: {
      id: number; 
      name: string;
      type: string;
      passDate: string;
      rePassDate: string;
      hasProtocol: string;
    }[];
  }

export type Training = {
    id: number;
    name: string;
    passDate: string;
    rePassDate: string;
    hasProtocol: string;
  };
  