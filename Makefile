.PHONY: examples

examples:
	+$(MAKE) -C examples/html/document_parse
	+$(MAKE) -C examples/html/document_parse_chunk
	+$(MAKE) -C examples/html/document_title
	+$(MAKE) -C examples/html/element_attributes
	+$(MAKE) -C examples/html/element_create
	+$(MAKE) -C examples/html/element_inner_html
	+$(MAKE) -C examples/html/elements_by_attr
	+$(MAKE) -C examples/html/elements_by_class_name
