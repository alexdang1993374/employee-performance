export interface IEmployee {
  id: number;
  name: string;
  performance: number;
  date: string;
}

export interface IEmployeeResponse {
  data: {
    data: {
      data: IEmployee[];
    };
  };
}
