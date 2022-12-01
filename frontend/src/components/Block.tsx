import { useBlock } from "../hooks/useBlock";
import { Textarea } from "./Input";
import { Box, Col } from "./Layout";
import { Button } from "./Button";

export function Block({ id }: any) {
  const { block, createChild, update } = useBlock(id);

  return (
    <Col align="stretch" css={{ marginTop: "$2" }}>
      <Box variant="bordered">
        <Textarea
          style={{ width: "100%" }}
          value={block?.value ?? ""}
          onChange={(e) => {
            update(e.target.value);
          }}
        />
      </Box>
      <Col align="stretch" css={{ spacingY: "$4", marginLeft: "$2" }}>
        {(block?.children?.length ?? 0) > 0
          ? block?.children.map((childBlock) => (
              <Block key={childBlock.id} id={childBlock.id} />
            ))
          : null}
      </Col>
      <Col css={{ marginBottom: "$4", marginTop: "-35px" }}>
        <Button
          onClick={() => {
            createChild();
          }}
        >
          Add child
        </Button>
      </Col>
    </Col>
  );
}
