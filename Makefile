.PHONY: examples

examples:
	+$(MAKE) -C examples/html/document_parse
	+$(MAKE) -C examples/html/document_parse_chunk
	+$(MAKE) -C examples/html/document_title
	+$(MAKE) -C examples/html/element_attributes
