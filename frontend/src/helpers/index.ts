import { IEmployee, SortOptions } from "@/types";

export const formatDate = (dateString: string): string => {
  const date = new Date(dateString);

  const day = String(date.getDate()).padStart(2, "0");
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const year = date.getFullYear();

  return `${month}/${day}/${year}`;
};

export const sortEmployees = (
  employees: IEmployee[],
  sortOption: SortOptions
): IEmployee[] => {
  const sortedEmployees = [...employees];

  sortedEmployees.sort((a, b) => {
    switch (sortOption) {
      case "Name":
        return a.name.localeCompare(b.name);
      case "Performance H->L":
        return b.performance - a.performance;
      case "Performance L->H":
        return a.performance - b.performance;
      case "Date":
        return new Date(a.Date).getTime() - new Date(b.Date).getTime();
      default:
        return 0;
    }
  });

  return sortedEmployees;
};
