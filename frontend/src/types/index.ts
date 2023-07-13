export interface IEmployee {
  id: number;
  name: string;
  performance: number;
  Date: string;
}

export interface IEmployeeResponse {
  data: {
    data: {
      data: IEmployee[];
    };
  };
}

export interface MenuItem {
  label: string;
  onClick: VoidFunction;
}

export type SortOptions = "Name" | "Performance" | "Date";
