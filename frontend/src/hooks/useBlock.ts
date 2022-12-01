import { useEffect, useState } from "react";

const baseUrl = "http://localhost:8090/blocks";

export interface Block {
  id: string;
  value: any;
  children: Block[];
}

export function useBlock(id: any): {
  block: Block | null;
  update: (value: any) => Promise<void>;
  createChild: () => Promise<void>;
} {
  const [block, setBlock] = useState<Block | null>(null);

  let urlSuffix = "";

  if (id) {
    urlSuffix = `/${id}`;
  }

  useEffect(() => {
    fetch(baseUrl + urlSuffix)
      .then((res) => res.json())
      .then((data) => {
        setBlock(data);
        // setValue(data.value);
      });
  }, [urlSuffix, setBlock]);

  return {
    block,

    async update(value) {
      const res = await fetch(baseUrl + urlSuffix, {
        method: "PUT",
        body: value,
      });
      const data = await res.json();
      setBlock(data);
    },

    async createChild() {
      const res = await fetch(baseUrl + urlSuffix, {
        method: "POST",
        body: "",
      });
      const data = await res.json();

      if (block) {
        setBlock({ ...block, children: [...block.children, data] });
      }
    },
  };
}
