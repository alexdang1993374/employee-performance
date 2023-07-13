"use client";

import axios from "axios";
import { useEffect, useState } from "react";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { formatDate, sortEmployees } from "@/helpers";
import { IEmployee, IEmployeeResponse, MenuItem, SortOptions } from "@/types";
import Dropdown from "./ui/Dropdown";

const EmployeeTable = () => {
  const [employees, setEmployees] = useState<IEmployee[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [currentSort, setCurrentSort] =
    useState<SortOptions>("Performance H->L");

  const sortOptions: MenuItem[] = [
    {
      label: "Performance H->L",
      onClick: () => setCurrentSort("Performance H->L"),
    },
    {
      label: "Performance L->H",
      onClick: () => setCurrentSort("Performance L->H"),
    },
    {
      label: "Name",
      onClick: () => setCurrentSort("Name"),
    },
    {
      label: "Date",
      onClick: () => setCurrentSort("Date"),
    },
  ];

  useEffect(() => {
    const fetchData = async () => {
      try {
        const result: IEmployeeResponse = await axios.get(
          "http://localhost:5001/employees"
        );

        setEmployees(sortEmployees(result.data.data.data, "Performance H->L"));
        setLoading(false);
      } catch (error) {
        console.error("Error fetching data: ", error);
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  useEffect(() => {
    setEmployees(sortEmployees(employees, currentSort));
  }, [currentSort]);

  return (
    <>
      <div className="mb-4 flex flex-col gap-4 items-center">
        <p className="text-lg">Employee Table</p>

        <div className="flex gap-2 items-center">
          <p>Sort by</p>

          <Dropdown menuItems={sortOptions} buttonLabel={currentSort} />
        </div>
      </div>

      {loading ? (
        <div>Loading...</div>
      ) : (
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Employee ID</TableHead>
              <TableHead>Name</TableHead>
              <TableHead>Performance</TableHead>
              <TableHead className="text-right">Date</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {employees.map((employee) => (
              <TableRow key={employee.id}>
                <TableCell className="font-medium">{employee.id}</TableCell>
                <TableCell>{employee.name}</TableCell>
                <TableCell>{employee.performance}</TableCell>
                <TableCell className="text-right">
                  {formatDate(employee.Date)}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </>
  );
};

export default EmployeeTable;
