import EmployeeTable from "@/components/EmployeeTable";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center p-24 md:p-0">
      <EmployeeTable />
    </main>
  );
}
