import EmployeeTable from "@/components/EmployeeTable";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <EmployeeTable />
    </main>
  );
}
